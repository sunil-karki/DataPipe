import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { FileManagerComponent } from './fileManager/fileManager.component';
import { InventorysComponent } from './inventory/inventorys.component';

const routes: Routes = [
  { path: '', component: InventorysComponent },
  { path: 'filemanager', component: FileManagerComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
