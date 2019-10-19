import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl } from '@angular/forms';

@Component({
  selector: 'app-tasks',
  templateUrl: './tasks.component.html',
  styleUrls: ['./tasks.component.css']
})
export class TasksComponent implements OnInit {
  newTaskForm: FormGroup;
  taskName: string;
  taskMemo: string;
  taskQuadrant: string;
  quadrants: any;

  constructor() { }

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
    this.taskName = this.newTaskForm.get('name').value;
    this.taskMemo = this.newTaskForm.get('memo').value;
    this.taskQuadrant = this.newTaskForm.get('quadrant').value;

    console.log(this.taskName);
    console.log(this.taskMemo);
    console.log(this.taskQuadrant);
  }
}
