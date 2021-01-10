import { Component, OnInit } from '@angular/core';
import { DataTableComponent } from './dataTable/dataTable.component';
import { InventoryService } from './inventory.service';

import { Inventorys } from './inventorys.model';

@Component({
  selector: 'app-inventorys',
  templateUrl: './inventorys.component.html',
  styleUrls: ['./inventorys.component.css']
})
export class InventorysComponent implements OnInit {
  fileData: Inventorys[];

  constructor(private InventorysService: InventoryService) { }

  ngOnInit() {
    return this.InventorysService.getData()
                .subscribe(data => {
                  this.fileData = data;
                  console.log('test haah')
                  console.log(data)
                })
  }

}
