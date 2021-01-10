import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { AngularFileUploaderModule } from "angular-file-uploader";

import { AppComponent } from './app.component';
import { AppRoutingModule } from './app-routing.module';
import { HeaderComponent } from './header/header.component';
import { InventorysComponent } from './inventory/inventorys.component';
import { DataTableComponent } from './inventory/dataTable/dataTable.component';
import { FileManagerComponent } from './fileManager/fileManager.component';
import { FileUploadComponent } from './fileManager/fileUplaod/fileUpload.component';
import { DropdownDirective } from './shared/dropdown.directive';
// import { AgGridModule } from 'ag-grid-angular';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import { MatTableModule } from '@angular/material/table';
import { MatPaginatorModule } from '@angular/material/paginator';
import { MatSortModule } from '@angular/material/sort';
import { InventoryService } from './inventory/inventory.service';

const materialModules = [
  MatTableModule,
  MatPaginatorModule,
  MatSortModule
];

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    InventorysComponent,
    DataTableComponent,
    FileManagerComponent,
    FileUploadComponent,
    DropdownDirective
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpClientModule,
    AngularFileUploaderModule,
    // AgGridModule.withComponents([]),
    materialModules,

    AppRoutingModule,
    NoopAnimationsModule
  ],
  exports: [
    materialModules
  ],
  providers: [InventoryService],
  bootstrap: [AppComponent]
})
export class AppModule { }
