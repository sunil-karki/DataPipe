import { Injectable } from '@angular/core';
// import {Http} from '@angular/http';
// import {environment} from '.../environments/environment';
// import 'rxjs/add/operator/map';
import { HttpClient } from '@angular/common/http';
import { Inventorys } from './inventorys.model'

@Injectable()
export class InventoryService {

  apiUrl = 'http://localhost:9090/';
//   apiUrl = 'http://mysafeinfo.com/api/data?list=englishmonarchs&format=json';

  constructor(private http: HttpClient) { }

  getData() {
    // return this.http.get(`${environment.serverUrl}/hello-world`)
    //   .map(response => response.json());
    return this.http.get<Inventorys[]>(this.apiUrl);
  }

}