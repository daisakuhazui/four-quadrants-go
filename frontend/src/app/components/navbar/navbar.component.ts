import { Component, Input, OnInit } from '@angular/core';
import { AuthUseCaseService } from '../../domains/auth/auth-use-case.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  @Input()
  displayUserName: string;

  constructor(public auth: AuthUseCaseService) { }

  ngOnInit() {
  }

}
