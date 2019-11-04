import { Component, OnInit } from '@angular/core';
import { Task } from '../models/task';
import { HttpClient } from '@angular/common/http';
import {HttpClientService } from '../service/http-client.service';

@Component({
  selector: 'app-tasks-index',
  templateUrl: './tasks-index.component.html',
  styleUrls: ['./tasks-index.component.css']
})
export class TasksIndexComponent implements OnInit {
  resTasks: Task[];
  firstQuadrantTasks: Task[];
  secondQuadrantTasks: Task[];
  thirdQuadrantTasks: Task[];
  fourthQuadrantTasks: Task[];

  constructor(
    private http: HttpClient,
    private httpClientService: HttpClientService
  ) { }

  ngOnInit() {
    this.getTasks();
  }

  // タスク一覧を取得する処理
  // TODO: ログインユーザーのタスクのみ取得できるようにする
  getTasks() {
    // ヘッダ情報セット
    const requestUri = this.httpClientService.host + '/tasks';
    // this.httpClientService.httpOptions = this.httpClientService.httpOptions.set('Access-Control-Allow-Origin', requestUri);

    // API 実行
    this.http.get(requestUri, this.httpClientService.httpOptions)
      .toPromise()
      .then((res) => {
        const response: any = res;
        this.resTasks = response;
      })
      .catch(
        // TODO: ここにエラーハンドリングを書く
      );
  }

}
