# Digital Commons BePress API

Based off the API documentation found here: https://bepress.com/wp-content/uploads/2020/05/Real-time-API-documentation.pdf

## Basics
- digital commons is accessed using the content-out.digitalcommons.com/v2 url. 
- the url is appended with the 'site_url' specific to each university or library
- then the url is appended with one of four endpoints:

1. query GET
2. export PUT
3. download GET
4. fields GET

## How to use
From your the GO code, run:
`go get https://github.com/AvalLegal/bepress-digital-commons-go-sdk/`

Then in the code, 
