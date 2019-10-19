import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl } from '@angular/forms';
import { Task } from '../models/task';
import { HttpClient } from '@angular/common/http';
import {HttpClientService } from '../service/http-client.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-tasks',
  templateUrl: './tasks.component.html',
  styleUrls: ['./tasks.component.css']
})
export class TasksComponent implements OnInit {
  newTaskForm: FormGroup;
  quadrants: any;
  task: Task;

  constructor(
    private http: HttpClient,
    private httpClientService: HttpClientService,
    private router: Router
  ) { }

  ngOnInit() {
    this.newTaskForm = new FormGroup({
      name: new FormControl(),
      memo: new FormControl(),
      quadrant: new FormControl(),
    });
    this.quadrants = ['第1象限', '第2象限', '第3象限', '第4象限'];
  }

  // 登録ボタン押下時の処理
  onSubmmit() {
    // フォーム入力内容をマッピング
    this.task = {
      id: 0,
      name: this.newTaskForm.get('name').value,
      memo: this.newTaskForm.get('memo').value,
      quadrant: this.newTaskForm.get('quadrant').value,
      completeFlag: false,
    };

    // ヘッダ情報
    const requestUri = this.httpClientService.host + '/task';
    this.httpClientService.httpOptions = this.httpClientService.httpOptions.headers.set('Access-Control-Allow-Origin', requestUri);

    console.log(this.httpClientService.httpOptions.headers);

    // API 実行
    this.http.post(requestUri, this.task, this.httpClientService.httpOptions)
      .subscribe(
        (res) => {
          // const response: any = res;
          // return response;
          this.router.navigate(['/']);
        },
        // TODO: エラーハンドリングを別途共通化して実装する
        (error) => {
          console.log(error);
        }
      );
  }
}
