
# **goparse**
This project is a solution to the **html/template** package that does not allow us to parse files into each subdirectory within it, as shown in the tree below.

```
views
├── components
│   ├── header.html
│   └── subcomponents
│       └── menu_header.html
└── index.html
```
with goparse, all the files inside will be placed in a temporary folder named **.goparse-tmp** which will later be used as a template

# **how it works**
![Shot-2024-12-15-083217](https://github.com/user-attachments/assets/b5962276-da08-42b5-9b2c-9b60d26cf47c)