import { Component, OnInit } from '@angular/core';
import { AuthUseCaseService } from './domains/auth/auth-use-case.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'Four Quadrants Go';

  constructor(private auth: AuthUseCaseService) {
  }

  ngOnInit(): void {
    this.auth.localAuthSetup();
  }
}
