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
  updatedTask: Task;
  firstQuadrantTasks: Task[];
  secondQuadrantTasks: Task[];
  thirdQuadrantTasks: Task[];
  fourthQuadrantTasks: Task[];
  isOpen: boolean;
  focusTaskId: number;

  constructor(
    private http: HttpClient,
    private httpClientService: HttpClientService
  ) { }

  ngOnInit() {
    this.getTasks();
    this.focusTaskId = 0;
    this.isOpen = false;
  }

  // タスク一覧を取得する処理
  // TODO: ログインユーザーのタスクのみ取得できるようにする
  getTasks() {
    // ヘッダ情報セット
    const requestUri = this.httpClientService.host + '/tasks';
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

  // タスクを完了する処理
  checkTask(taskId: number) {
    // ヘッダ情報セット
    const requestUri = this.httpClientService.host + '/task/check/' + taskId.toString();
    // API 実行
    this.http.put(requestUri, this.httpClientService.httpOptions)
      .toPromise()
      .then((res) => {
        const response: any = res;
        this.updatedTask = response;
        // TODO: 完了扱いのタスクは文字色が薄くなるようにスタイルをつける
        return this.updatedTask;
      })
      .catch(
        // TODO: ここにエラーハンドリングを書く
      );
  }

  // タスクを編集する処理
  editTask(task: Task) {
    this.focusTaskId = 0;
    this.isOpen = false;

    // ヘッダ情報セット
    const requestUri = this.httpClientService.host + '/task/' + task.id.toString();
    // API 実行
    this.http.put(requestUri, task, this.httpClientService.httpOptions)
      .toPromise()
      .then((res) => {
        const response: any = res;
        this.updatedTask = response;
        return this.updatedTask;
      })
      .catch(
        // TODO: ここにエラーハンドリングを書く
      );
  }

  // isOpen フラグを切り替える
  swich(taskId: number) {
    this.focusTaskId = taskId;
    this.isOpen = true;
  }
}
