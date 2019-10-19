import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { TasksIndexComponent } from './tasks-index/tasks-index.component';
import { TasksComponent } from './tasks/tasks.component';

const routes: Routes = [
  {path: '', component: TasksIndexComponent},
  {path: 'tasks', component: TasksComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
