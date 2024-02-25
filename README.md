# GetDocxText

### 中文版

```markdown
# GetDocxText

## 简介

`GetDocxText`是一个Go语言编写的库，用于从Microsoft Word文档（.docx格式）中提取文本内容。该库提供了两个主要的公共方法，使开发者能够轻松集成此功能到他们的项目中。

## 功能

- `GetXmlContent(docxFile string) ([]byte, error)`: 从.docx文件中提取XML内容。
- `GetTextByParagraph(content []byte) ([]string, error)`: 解析XML内容，按段落提取文本。

## 安装

```bash
go get github.com/Art-Man/GetDocxText
```

## 使用示例

```go
package main

import (
    "github.com/Art-Man/GetDocxText/docxtext"
    "log"
)

func main() {
    xmlContent, err := docxtext.GetXmlContent("./testdata/example.docx")
    if err != nil {
        log.Fatalf("Error extracting XML content: %v", err)
    }

    texts, err := docxtext.GetTextByParagraph(xmlContent)
    if err != nil {
        log.Fatalf("Error getting text by paragraph: %v", err)
    }

    for _, text := range texts {
        log.Println(text)
    }
}
```

## 测试

运行以下命令以执行单元测试：

```bash
go test -v ./docxtext
```

## 贡献

欢迎通过GitHub pull requests和issues参与贡献。

## 许可证

该项目使用MIT许可证，详情请见LICENSE文件。
```

### English Version

```markdown
# GetDocxText

## Introduction

`GetDocxText` is a Go library for extracting text content from Microsoft Word documents (.docx format). It provides two main public methods, enabling developers to easily integrate this functionality into their projects.

## Features

- `GetXmlContent(docxFile string) ([]byte, error)`: Extracts XML content from a .docx file.
- `GetTextByParagraph(content []byte) ([]string, error)`: Parses the XML content and extracts text by paragraph.

## Installation

```bash
go get github.com/Art-Man/GetDocxText
```

## Usage Example

```go
package main

import (
    "github.com/Art-Man/GetDocxText/docxtext"
    "log"
)

func main() {
    xmlContent, err := docxtext.GetXmlContent("./testdata/example.docx")
    if err != nil {
        log.Fatalf("Error extracting XML content: %v", err)
    }

    texts, err := docxtext.GetTextByParagraph(xmlContent)
    if err != nil {
        log.Fatalf("Error getting text by paragraph: %v", err)
    }

    for _, text := range texts {
        log.Println(text)
    }
}
```

## Testing

Run the following command to execute unit tests:

```bash
go test -v ./docxtext
```

## Contributing

Contributions via GitHub pull requests and issues are welcome.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
```

这些README文件提供了关于如何安装、使用和为`GetDocxText`项目贡献的基本指南。确保在项目的根目录下创建`README.md`文件，并将上述内容粘贴到相应的文件中。