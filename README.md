# Introduction

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
- [ ] Join 2 PDFs
- [ ] Convert image to PDF
- [ ] PDF Checking
- [ ] Extract PDF content
- [ ] Compress PDF
- [ ] Reorder PDF pages
- [ ] Combine multiple PDF

## Example

1. Get PDF file info

```
http://localhost:1323/api/v1/pdf/info?path=/storage/pdf&name=camry_ebrochure.pdf
```

2. Split PDF

```
curl -X POST http://localhost:1323/api/v1/pdf/split \
-H 'Content-Type: application/json' \
-d '{"name":"camry_ebrochure.pdf","path":"/storage/pdf", "range":"1-3,3-4"}'
```

3. Join 2 PDFs



