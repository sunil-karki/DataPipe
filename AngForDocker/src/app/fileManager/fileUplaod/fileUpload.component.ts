import { HttpClient, HttpEventType } from '@angular/common/http';
import { Component } from '@angular/core';


@Component({
  selector: 'app-fileUpload',
  templateUrl: './fileUpload.component.html',
  styleUrls: []
})
export class FileUploadComponent {
    // set null to interface value by -->   {} as any or  <File>{}
    selectedFile: File = <File>{};

    constructor(private http: HttpClient) {
       
    }

    onFileSelected(event: any) {
        this.selectedFile = <File>event.target.files[0];
    }

    onUpload() {
        const fd = new FormData();
        fd.append('file', this.selectedFile, this.selectedFile.name);
        this.http.post('http://localhost:9090/upload', fd, {
                reportProgress: true,
                observe: 'events'
              })
            .subscribe(event => {
                if(event.type == HttpEventType.UploadProgress) {
                    // console.log('Upload progress: ' + Math.round(event.loaded / event.total * 100) + '%')
                    console.log("Uploading ...")
                }
                else if(event.type == HttpEventType.Response) {
                    console.log(event)
                }
                console.log(event);
              })
    }
}