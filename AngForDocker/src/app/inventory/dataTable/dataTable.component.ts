import { Component, OnInit, ViewChild } from '@angular/core';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { InventoryService } from '../inventory.service';


@Component({
  selector: 'app-dataTable',
  templateUrl: './dataTable.component.html',
  styleUrls: ['./dataTable.component.css']
})
export class DataTableComponent implements OnInit {

  displayedColumns: string[] = ['position', 'fileid', 'filename', 'description', 'filedate', 'source'];
//   dataSource = ELEMENT_DATA;
  dataSource = new MatTableDataSource<PeriodicElement>();
  data : any = null;

  @ViewChild(MatPaginator) paginator: MatPaginator;

  ngAfterViewInit() {
    this.dataSource.paginator = this.paginator;
  }

  constructor(private InventorysService: InventoryService) {
    //   this.paginator = any;
   }

  ngOnInit() {
    // this.dataSource.paginator = this.paginator;
    return this.InventorysService.getData()
                .subscribe(data => {
                  this.data = data;
                  console.log(data);
                  this.dataSource = new MatTableDataSource(data);
                })
  }

}

/* Static data */ 

export interface PeriodicElement {
    Fileid: number;
    Position: number;
    Filename: string;
    Description: string;
    Filedate: string;
    Source: string;
}
  
// const ELEMENT_DATA: PeriodicElement[] = [
//     { position: 1, id: 1.0079, filename: 'Hydrogen', description: 'Sth abt File', filedate: '2020-02-11',  source: 'H' },
//     { position: 2, id: 1.0079, filename: 'Helium', description: 'Sth abt File', filedate: '2020-02-11', source: 'He' },
//     { position: 3, id: 1.0079, filename: 'Lithium', description: 'Sth abt File', filedate: '2020-02-11', source: 'Li' },
//     { position: 4, id: 1.0079, filename: 'Beryllium', description: 'Sth abt File', filedate: '2020-02-11', source: 'Be' },
//     { position: 5, id: 1.0079, filename: 'Boron', description: 'Sth abt File', filedate: '2020-02-11', source: 'B' },
//     { position: 6, id: 1.0079, filename: 'Carbon', description: 'Sth abt File', filedate: '2020-02-11', source: 'C' },
//     { position: 7, id: 1.0079, filename: 'Nitrogen', description: 'Sth abt File', filedate: '2020-02-11', source: 'N' },
//     { position: 8, id: 1.0079, filename: 'Oxygen', description: 'Sth abt File', filedate: '2020-02-11', source: 'O' },
//     { position: 9, id: 1.0079, filename: 'Fluorine', description: 'Sth abt File', filedate: '2020-02-11', source: 'F' },
//     { position: 10, id: 1.0079, filename: 'Neon', description: 'Sth abt File', filedate: '2020-02-11', source: 'Ne' },
// ];

// [{"id":1,"filename":"Latte","description":"Frothy milky coffee","filedate":"2.45","sku":"abc323"},{"id":2,"name":"Espresso","description":"Short and strong coffee without milk","price":1.99,"sku":"fjd34"}]
