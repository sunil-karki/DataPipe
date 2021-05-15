# Angular Project setup

**Step 1: Adding Bootstrap**
-------------------------------------------------------------------

  **_Option: 1 -  Add the three links on `index.html` as below_**
  
  [ Inside HEAD ] 
  ```html
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
  ```
  [ Inside Body ] 
  ```html
  <script src="https://code.jquery.com/jquery-3.1.1.min.js" integrity="sha256-hVVnYaiADRTO2PzUGmuLJr8BLUSjGIZsDYGmIJLv2b8=" crossorigin="anonymous"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script> -->
  ```

  **_Option: 2 -  Next Method_**  

  Run the following command:

    npm install bootstrap@3 jquery --save

  Then add the following in the `angular.json` file
  
    "styles": [
      "src/styles.css",
      "../node_modules/bootstrap/dist/css/bootstrap.min.css"
    ],
    "scripts": [
      "../node_modules/jquery/dist/jquery.min.js",
      "../node_modules/bootstrap/dist/js/bootstrap.min.js"
    ]

--------------------------------------------------------------------------------------------------------------------------------
--------------------------------------------------------------------------------------------------------------------------------            

**Step 2: Add Angular Material for DataTable**
------------------------------------------------------------------

    ng add @angular/material

Then add MatTableModule, MatPaginatorModule, MatSortModule to the `app.module.ts` file.

--------------------------------------------------------------------------------------------------------------------------------
--------------------------------------------------------------------------------------------------------------------------------


# Sample

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 11.0.6.

## Development server

Run `ng serve` for a dev server. Navigate to `http://localhost:4200/`. The app will automatically reload if you change any of the source files.

## Code scaffolding

Run `ng generate component component-name` to generate a new component. You can also use `ng generate directive|pipe|service|class|guard|interface|enum|module`.

## Build

Run `ng build` to build the project. The build artifacts will be stored in the `dist/` directory. Use the `--prod` flag for a production build.

## Running unit tests

Run `ng test` to execute the unit tests via [Karma](https://karma-runner.github.io).

## Running end-to-end tests

Run `ng e2e` to execute the end-to-end tests via [Protractor](http://www.protractortest.org/).

## Further help

To get more help on the Angular CLI use `ng help` or go check out the [Angular CLI Overview and Command Reference](https://angular.io/cli) page.
