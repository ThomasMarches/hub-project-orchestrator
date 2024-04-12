import { Component } from '@angular/core';
import { ApiService } from './services/api.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent {
  title = 'my-app';

  constructor(private dataService: ApiService) {
    this.dataService.fetchData('/').subscribe((data) => {
      console.log(data);
    });
  }
}
