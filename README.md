# GoPDF

[![codecov][codecov-badge]][codecov-link]

## Introduction

Our customers upload documents and we make sure we Print them. Besides that we also provide
tooling to work with PDF documents, usually in the form of an API. That is where your assignment
comes in.

## Specification

- HTTP Protocols
- Integration Test
- Echo Framework

## Features

- [x] show PDF description detail
- [x] split PDF
- [x] Combine multiple PDF
- [x] Convert JPG to PDF
- [ ] PDF Checking
- [ ] Extract PDF content
- [x] Compress PDF
- [x] Reorder PDF pages


## How to Use


1. Get PDF file info

```
http://localhost:1323/api/v1/pdf/info?path=/storage/pdf&name=camry_ebrochure.pdf
```

2. Split PDF

```
curl -X POST http://localhost:1323/api/v1/pdf/split \
-H 'Content-Type: application/json' \
-d '{"name":"camry_ebrochure.pdf","path":"/storage/pdf", "range":"1-3,5,8-9"}'
```



3. Merge PDFs

```
curl -X POST http://localhost:1323/api/v1/pdf/merge \
-H 'Content-Type: application/json' \
-d '{"infiles":[{"name":"camry_ebrochure.pdf","path":"/storage/pdf"},{"name":"mirai_ebrochure.pdf","path":"/storage/pdf"}], "outfile":"camry_mirai_ebrochure.pdf"}'
```

4. Convert JPG to PDF

```
curl -X POST http://localhost:1323/api/v1/pdf/jpg-to-pdf \
-H 'Content-Type: application/json' \
-d '{"infiles":[{"name":"camry_ebrochure.pdf","path":"/storage/pdf"},{"name":"mirai_ebrochure.pdf","path":"/storage/pdf"}], "outfile":"camry_mirai_ebrochure.pdf"}'
```

5. Reorder PDF pages

```
curl -X POST http://localhost:1323/api/v1/pdf/reorder \
-H 'Content-Type: application/json' \
-d '{"name":"camry_ebrochure.pdf","path":"/storage/pdf", "page_number":"1-2", "new_page_number":"4"}'
```

6. Compress PDF

```
curl -X POST http://localhost:1323/api/v1/pdf/compress \
-H 'Content-Type: application/json' \
-d '{"name":"camry_ebrochure.pdf","path":"/storage/pdf", "page_number":"1-2", "new_page_number":"4"}'
```

## Credits

- [pdfcpu api](https://pkg.go.dev/github.com/hhrutter/pdfcpu@v0.2.2/pkg/api)
 

## References

- PDF Samples https://www.toyota.com/brochures/cars-minivan/
- Car Image by <a href="https://unsplash.com/@dhivakrishna?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Dhiva Krishna</a> on <a href="https://unsplash.com/photos/black-mercedes-benz-car-YApS6TjKJ9c?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Unsplash</a>
- Photo by <a href="https://unsplash.com/@dpanyukov?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Dima Panyukov</a> on <a href="https://unsplash.com/photos/DwxlhTvC16Q?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Unsplash</a>
- Photo by <a href="https://unsplash.com/@neonbrand?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Kenny Eliason</a> on <a href="https://unsplash.com/photos/FcyipqujfGg?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Unsplash</a>
  

[codecov-badge]: https://codecov.io/gh/aysf/gopdf/branch/master/graph/badge.svg?token=x
[codecov-link]: https://codecov.io/gh/aysf/gopdf
  