import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { TasksIndexComponent } from './tasks-index/tasks-index.component';
import { TasksComponent } from './tasks/tasks.component';
import { UserAttendanceComponent } from './components/user-attendance/user-attendance.component';
import { CallbackComponent } from './components/auth/callback/callback.component';

const routes: Routes = [
  {
    path: '',
    redirectTo: 'user-attendance',
    pathMatch: 'full'
  },
  {
    path: 'index',
    component: TasksIndexComponent
  },
  {
    path: 'tasks',
    component: TasksComponent
  },
  {
    path: 'user-attendance',
    component: UserAttendanceComponent
  },
  {
    path: 'callback',
    component: CallbackComponent
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
