import { Component, OnInit } from '@angular/core';
import { AuthUseCaseService } from '../../../domains/auth/auth-use-case.service';

@Component({
  selector: 'app-callback',
  templateUrl: './callback.component.html',
  styleUrls: ['./callback.component.css']
})
export class CallbackComponent implements OnInit {

  constructor(private auth: AuthUseCaseService) { }

  ngOnInit() {
    this.auth.handleAuthCallback();
  }

}
