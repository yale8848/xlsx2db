# Hyperlinks
[[toc]]

You can add any type of hyperlink to cell or range (row and col are ranges too), such as: hyperlink to file, network, website, email, another sheet and etc. 

Xlsx2Go supports string version of hyperlinks, as well as custom version via special type. Read about types below.

::: warning Hyperlinks limits
Excel has built-in limit for total number of hyperlinks per sheet. Check [Excel Limits](/guide/limits.md)
:::

::: tip Tip
In some cases, adding same hyperlink to range (instead of cells directly) can help to exceed these limits.
:::

::: note Filepath
Xlsx2Go automatically detects windows, unix and network versions of filepath and validates it.
:::

### Hyperlinks and cells

```go
	//to add hyperlink to cell
	sheet.CellByRef("N28").SetHyperlink("http://google.com")

	//to add hyperlink to cell and value same time
	sheet.CellByRef("N28").SetValueWithHyperlink(12345, "http://google.com")

	//to remove hyperlink
	sheet.CellByRef("N28").RemoveHyperlink()

	//to retrieve hyperlink information
	link := sheet.CellByRef("N28").Hyperlink()
```

### Hyperlinks and ranges
```go
	//add hyperlink to range
	sheet.RangeByRef("A1:A10").SetHyperlink("http://google.com")

	//to remove hyperlink from range
	sheet.RangeByRef("A1:A10").RemoveHyperlink()

	//Row and Column are subtypes of Range, so inherit functionality
	sheet.Row(0).SetHyperlink("http://google.com")
	sheet.Col(0).SetHyperlink("http://google.com")
```

::: note N.B.:
In case of ranges, only one hyperlink will be added to sheet, but for whole range!
:::

### String hyperlink
As was shown, the simplest way to add hyperlink is to use string version. At the same time, that version is quite limited - valid format should be used and you can't set tooltips, styles and etc. 
:::tip 
Xlsx2Go supports same format as Excel's `HYPERLINK` function and will try to autodetect type of hyperlink (email, file, website and etc) to validate it later. 

Check [Excel Hyperlink]( https://support.office.com/en-us/article/hyperlink-function-333c7ce6-c5ae-4164-9c47-7de9b76f577f) for more information.
:::

::: note Same file, same sheet
* `#A1`  
:::

::: note Same file, other sheet
* `#SheetName!A1`  
* `#'Sheet Name'!A1`  
:::

::: note External file
* `D:\Folder\File.docx`
* `[D:\Folder\File.docx]`  
* `D:\Folder\File.docx#Bookmark`  
* `[D:\Folder\File.docx]Bookmark`  
* `D:\Folder\File.xlsx#SheetName!A1`  
* `[D:\Folder\File.xlsx]SheetName!A1`  
* `D:\Folder\File.xlsx#'Sheet Name'!A1`  
* `[D:\Folder\File.xlsx]'Sheet Name'!A1`  
:::

::: note Network file
* `\\SERVER\Folder\File.doc`  
* `\\SERVER\Folder\File.xlsx#SheetName!A1`  
* `[\\SERVER\Folder\File.xlsx]SheetName!A1`  
* `\\SERVER\Folder\File.xlsx#'Sheet Name'!A1`  
* `[\\SERVER\Folder\File.xlsx]'Sheet Name'!A1`  
:::

::: note Url
* `https://www.spam.it`  
* `https://www.spam.it/#bookmark`  
* `[https://www.spam.it/]bookmark`  
:::

::: note Email
* `mailto:spam@spam.it`  
* `mailto:spam@spam.it?subject=topic`  
:::

### Custom hyperlink
While with string version of hyperlink you can add any kind of hyperlink, sometimes we need additional settings like tooltips or styles. For these cases you can use special type and configure hyperlink as you wish.

#### Example

<<< @/src/code/hyperlinks_test.go

::: tip
As was shown in example, few relevant targets can be combined:
* `ToFile` + `ToBookmark`, e.g. Word file with bookmark or Excel file with named region
* `ToFile` + `ToRef`, e.g. reference at other Excel file
* `ToUrl` + `ToBookmark`, e.g. bookmark at other site
::: 

#### Target types + examples
```go
	//sets target to external file
	// can be unix, windows or network format
	hyperlink.ToFile(`./example_simple.xlsx`)

	//sets target to web site
	hyperlink.ToUrl(`https://www.spam.it`)

	//sets target to ref of sheet with sheetName 
	// omit sheetName to set location to ref of active sheet
	hyperlink.ToRef(`C3`, "Sheet1")

	//sets target to bookmark 
	// can be named region in xlsx, bookmark of remote file or even site
	hyperlink.ToBookmark(`bookmark`)

	//sets target to email
	hyperlink.ToMail(`spam@spam.it`, "topic")

	//sets target in same format as string version of hyperlink
	hyperlink.ToTarget(`D:\Folder\File.docx#Bookmark`)
```
