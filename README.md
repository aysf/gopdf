# GoPDF

**

<!-- [![codecov][codecov-badge]][codecov-link] -->
[![codecov](https://codecov.io/gh/aysf/gopdf/graph/badge.svg?token=2QB9QJ2V7P)](https://codecov.io/gh/aysf/gopdf)

## Introduction

Providing tooling to work with PDF documents, usually in the form of an API. 

## Specification

- HTTP Protocols
- Integration Test
- Echo Framework

## Features

- [x] show PDF description detail
- [x] split PDF
- [x] trim PDF
- [x] remove PDF
- [x] merge PDF (combine multiple PDF)
- [x] convert JPG to PDF
- [x] compress PDF
- [x] reorder PDF pages


## How to Use

### Command

- to run the API use `go run cmd/api/*.go` or `make api`
- to test run `go test -v ./...` or `make test`
- to see test coverage, run `test makec` and `go tool cover -html=coverage.out`

### Example Request

1. Get PDF file info

This example request shows information about the pdf file

```
http://localhost:1323/api/v1/pdf/info?path=/storage/testPdf&name=camry_ebrochure.pdf
```

2. Split PDF

This example request generates single or multiples splitted pdf files based on `selected_pages` field input

```
curl -X POST http://localhost:1323/api/v1/pdf/split \
-H 'Content-Type: application/json' \
-d '{
    "name": "yle-flyers-sample.pdf",
    "path": "/storage/testPdf",
    "selected_pages": ["3-29","1-2","32-36"]
}'
```

3. Trim PDF

This example request generates a single pdf file that contains all of the targeted files.

```
curl -X POST http://localhost:1323/api/v1/pdf/trim \
-H 'Content-Type: application/json' \
-d '{
    "inname": "yle-flyers-sample.pdf",
    "inpath": "/storage/testPdf",
    "outname": "yle-flyers-sample_trimmed.pdf",
    "outpath": "/storage/testPdf",
    "target_pages": ["1-3","8-9"]
}'
```

4. Remove PDF

This example request generates a new copy of the pdf with the removed pages listed on 'target_pages' field

```
curl -X POST http://localhost:1323/api/v1/pdf/remove \
-H 'Content-Type: application/json' \
-d '{
    "inname": "yle-flyers-sample.pdf",
    "inpath": "/storage/testPdf",
    "outname": "yle-flyers-sample_removed.pdf",
    "outpath": "/storage/testPdf",
    "target_pages": ["1-3","5","8-9"]
}'
```

5. Merge PDFs

This example request merges two or more pdf files

```
curl -X POST http://localhost:1323/api/v1/pdf/merge \
-H 'Content-Type: application/json' \
-d '{
  "infiles": [
    {
      "name": "camry_ebrochure.pdf",
      "path": "/storage/testPdf"
    },
    {
      "name": "mirai_ebrochure.pdf",
      "path": "/storage/testPdf"
    }
  ],
  "outfile": "camry_mirai_ebrochure.pdf"
}'

```

5. Convert JPG to PDF

This example request creates a single pdf file from single or multiple image(s)

```
curl -X POST http://localhost:1323/api/v1/pdf/jpg-to-pdf \
-H 'Content-Type: application/json' \
-d '{
  "infiles": [
    {
      "name": "dhiva-krishna-YApS6TjKJ9c-unsplash.jpg",
      "path": "/storage/testImage"
    },
    {
      "name": "dima-panyukov-DwxlhTvC16Q-unsplash.jpg",
      "path": "/storage/testImage"
    },
    {
      "name": "kenny-eliason-FcyipqujfGg-unsplash.jpg",
      "path": "/storage/testImage"
    }
  ],
  "outfile": "jpg-to-pdf-output.pdf",
  "outpath": "/storage/testImage",
  "configs": {
    "page_size": "A4",
    "scale": 0.95
  }
}'

```

6. Compress PDF

This example request generates a new compressed pdf file from the original pdf

```
curl -X POST http://localhost:1323/api/v1/pdf/compress \
-H 'Content-Type: application/json' \
-d '{
    "infile": "camry_ebrochure.pdf",
    "inpath": "/storage/testPdf",
    "outfile": "camry_ebrochure_compressed.pdf",
    "outpath": "/storage/testPdf"
}'
```

7. Reorder PDF pages

This request example generates new pdf file containing pages that have been sorted based on `new_page_order` field input.

```
curl -X POST http://localhost:1323/api/v1/pdf/reorder \
-H 'Content-Type: application/json' \
-d '{
    "inname": "yle-flyers-sample.pdf",
    "inpath": "/storage/testPdf",
    "outname": "yle-flyers-sample_reordered.pdf",
    "outpath": "/storage/testPdf",
    "new_page_order": ["3-31","1-2","32-36"]
}'
```

## Credits

- It wraps the absolutely amazing [pdfcpu api](https://pkg.go.dev/github.com/hhrutter/pdfcpu@v0.2.2/pkg/api) library
- Heavily inspired by [ilovepdf](https://www.ilovepdf.com/), [smallpdf](https://smallpdf.com/pdf-tools), [pdf2go](https://www.pdf2go.com/), [pdf.io](https://pdf.io/), etc.
 

## References

- PDF Sample 1 by [Toyota](https://www.toyota.com/brochures/cars-minivan/)
- PDF Sample 2 by [Cambridge English](https://www.cambridgeenglish.org/latinamerica/images/165873-yle-sample-papers-flyers-vol-1.pdf)
- Car Photo 1 by <a href="https://unsplash.com/@dhivakrishna?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Dhiva Krishna</a> on <a href="https://unsplash.com/photos/black-mercedes-benz-car-YApS6TjKJ9c?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Unsplash</a>
- Car Photo 2 by <a href="https://unsplash.com/@dpanyukov?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Dima Panyukov</a> on <a href="https://unsplash.com/photos/DwxlhTvC16Q?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Unsplash</a>
- Car Photo 3 Photo by <a href="https://unsplash.com/@neonbrand?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Kenny Eliason</a> on <a href="https://unsplash.com/photos/FcyipqujfGg?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Unsplash</a>
  

<!-- [codecov-badge]: https://codecov.io/gh/aysf/gopdf/branch/master/graph/badge.svg?token=2QB9QJ2V7P
[codecov-link]: https://codecov.io/gh/aysf/gopdf -->
  
