import { Routes } from '@angular/router';
import { HomeComponent } from './components/home/home.component';
import { AboutComponent } from './components/about/about.component';
import { TestFormComponent } from './components/test-form/test-form.component';
import { CoursesComponent } from './components/courses/courses.component';
import { ThreadsComponent } from './components/threads/threads.component';

export const routes: Routes = [
  { path: '', redirectTo: 'home', pathMatch: 'full' },
  { path: 'home', component: HomeComponent },
  { path: 'threads', component: ThreadsComponent },
  { path: 'courses', component: CoursesComponent },
  { path: 'about', component: AboutComponent },
  { path: 'form', component: TestFormComponent }
];
