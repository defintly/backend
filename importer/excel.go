package importer

import (
	"fmt"
	"github.com/defintly/backend/database"
	"github.com/xuri/excelize/v2"
	"regexp"
	"strconv"
	"strings"
)

func ImportFromExcel(excelFileName string) {
	excel, err := excelize.OpenFile(excelFileName)
	if err != nil {
		panic(err)
	}

	excludeCharacterRegex, err := regexp.Compile("[^a-zA-Z0-9_]+")
	if err != nil {
		panic(err)
	}

	categoriesIndex := map[string]string{}
	collectionsIndex := map[string]string{}

	for _, sheetName := range []string{"Categories", "Criteria", "Collections", "Concepts"} {
		fmt.Printf("processing %s...\n", sheetName)

		unfilteredImportColumns, err := excel.GetCols(sheetName)
		if err != nil {
			panic(err)
		}

		var trimmedUnfilteredImportColumns [][]string

		emptyRows := 0

		for i := 0; i < len(unfilteredImportColumns[0]); i++ {
			if strings.TrimSpace(unfilteredImportColumns[0][i]) == "" {
				emptyRows++
			}

			// if more than 10 rows are empty, concat the empty rows (since the document might be finished then)
			if emptyRows > 10 {
				for j := 0; j < len(unfilteredImportColumns); j++ {
					var tempCol []string

					for k := 0; k <= i-emptyRows; k++ {
						tempCol = append(tempCol, unfilteredImportColumns[j][k])
					}

					if strings.TrimSpace(unfilteredImportColumns[j][0]) != "" {
						trimmedUnfilteredImportColumns = append(trimmedUnfilteredImportColumns, tempCol)
					}
				}

				break
			}
		}

		// if there are not enough empty rows, use the old table
		if emptyRows <= 10 {
			trimmedUnfilteredImportColumns = unfilteredImportColumns
		}

		var filteredImportColumns [][]string

		// give every row an id for relation
		idColumn := []string{"id"}
		for i := 1; i <= len(trimmedUnfilteredImportColumns[0])-1; i++ {
			intAsString := strconv.Itoa(i)
			idColumn = append(idColumn, intAsString)

			// map relational relevant data for later
			if sheetName == "Categories" {
				categoriesIndex[trimmedUnfilteredImportColumns[1][i]] = intAsString
			} else if sheetName == "Collections" {
				collectionsIndex[trimmedUnfilteredImportColumns[1][i]] = intAsString
			}
		}
		filteredImportColumns = append(filteredImportColumns, idColumn)

		var tempRelationalArray []string

		for _, row := range trimmedUnfilteredImportColumns {
			// filter redundant data
			if row[0] == "AGISI" || row[0] == "Contact" || row[0] == "Twitter" || row[0] == "Logo" ||
				row[0] == "Banner" || row[0] == "Related work" || row[0] == "AGISI more" || row[0] == "Cite App" ||
				row[0] == "Quality criteria in this category" ||
				row[0] == "Other quality criteria in the same category" || strings.TrimSpace(row[0]) == "" {
				continue
			}

			// filter even more redundant relational data by glide
			if sheetName == "Categories" {
				if row[0] == "Criteria=Criteria:Category:Multiple" {
					continue
				}
			} else if sheetName == "Collections" {
				if row[0] == "Concepts=Concepts:Concept:Multiple" {
					continue
				}
			} else if sheetName == "Criteria" || sheetName == "Concepts" {
				// save relations for later
				if row[0] == "Category=Categories:Category" {
					tempRelationalArray = row
					continue
				}

				if row[0] == "Concept" {
					tempRelationalArray = row
				}
			}

			// sanitizing (remove special characters etc.)
			row[0] = strings.ReplaceAll(strings.ToLower(row[0]), " ", "_")
			row[0] = excludeCharacterRegex.ReplaceAllString(row[0], "")

			filteredImportColumns = append(filteredImportColumns, row)
		}

		// add relations by foreign ids of tables
		if sheetName == "Criteria" {
			categoriesArray := []string{"category_id"}
			for i, categoryName := range tempRelationalArray {
				if i == 0 {
					continue
				}

				categoriesArray = append(categoriesArray, categoriesIndex[categoryName])
			}

			filteredImportColumns = append(filteredImportColumns, categoriesArray)
		}

		if sheetName == "Concepts" {
			collectionsArray := []string{"collection_id"}
			for i, collectionName := range tempRelationalArray {
				if i == 0 {
					continue
				}

				collectionsArray = append(collectionsArray, collectionsIndex[collectionName])
			}

			filteredImportColumns = append(filteredImportColumns, collectionsArray)
		}

		tableName := strings.ToLower(sheetName)

		// generate DDL
		createQuery := "CREATE TABLE IF NOT EXISTS " + tableName + "("
		for _, rowName := range filteredImportColumns {
			createQuery += "\"" + rowName[0] + "\""
			if strings.HasSuffix(rowName[0], "id") {
				if rowName[0] == "id" {
					createQuery += " SERIAL PRIMARY KEY, "
				} else {
					createQuery += " INTEGER, "
				}
			} else {
				createQuery += " TEXT, "
			}
		}
		createQuery = strings.TrimSuffix(createQuery, ", ") + ")"

		// generate INSERT statement
		insertQuery := "INSERT INTO " + tableName + " VALUES "
		for i := 1; i < len(filteredImportColumns[0]); i++ {
			insertQuery += "("
			for j := 0; j < len(filteredImportColumns); j++ {
				data := filteredImportColumns[j][i]

				data = strings.ReplaceAll(data, `
`, "\\n")
				// escape simple quotes
				data = strings.ReplaceAll(data, "'", "''")

				_, err := strconv.Atoi(data)
				if err == nil {
					insertQuery += data + ", "
				} else {
					insertQuery += "'" + data + "', "
				}

			}
			insertQuery = strings.TrimSuffix(insertQuery, ", ") + "), "
		}
		insertQuery = strings.TrimSuffix(insertQuery, ", ") + " ON CONFLICT DO NOTHING"

		database.MustExec(createQuery)
		database.MustExec(insertQuery)
	}
}
