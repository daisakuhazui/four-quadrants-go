import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { TopComponent } from './top/top.component';
import { TasksIndexComponent } from './tasks-index/tasks-index.component';
import { TasksComponent } from './tasks/tasks.component';
import { UserAttendanceComponent } from './components/user-attendance/user-attendance.component';
import { CallbackComponent } from './components/auth/callback/callback.component';
import { AuthGuard } from './domains/guard/auth.guard';

const routes: Routes = [
  {
    path: '',
    component: TopComponent,
  },
  {
    path: 'index',
    component: TasksIndexComponent,
    canActivate: [AuthGuard]
  },
  {
    path: 'tasks',
    component: TasksComponent,
    canActivate: [AuthGuard]
  },
  {
    path: 'user-attendance',
    component: UserAttendanceComponent,
    canActivate: [AuthGuard]
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
