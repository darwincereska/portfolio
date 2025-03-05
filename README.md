# Darwin Cereska Portfolio

## Overview
This is a portfolio website built with Astro, a modern static site builder, and Go for the backend and less resource consumption.

## Tech Stack
- Astro: Frontend framework
- Go: Backend server
- HTML/CSS/JavaScript 

## Prerequisites
- Node.js v16+
- Go v1.16+
- npm or yarn

## Installation

1. Clone the repository
```bash
git clone https://github.com/darwincereska/portfolio.git
cd portfolio
```

2. Install frontend dependencies
```bash
npm install
```

3. Build the Go backend
```bash 
go build -o server main.go
```

## Running the Project

1. Build the Astro project for production
```bash
npm run build
```

2. Run the Go backend server
```bash
./server
```

The site will be available at by default `http://localhost:9999`

## Running for development

1. Run astro in dev mode
``` bash
npm run dev
```