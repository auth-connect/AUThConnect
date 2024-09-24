import { Routes } from '@angular/router';

export const routes: Routes = [
  {
    path: '',
    loadComponent: () => import('./components/shared/layout/layout.component').then(m => m.LayoutComponent),
    children: [
        {
          path: 'home',
          loadComponent: () => import('./components/pages/home/home.component').then(m => m.HomeComponent),
          data: { title: 'Home' },
        },
        {
          path: 'threads',
          loadComponent: () => import('./components/pages/threads/threads.component').then(m => m.ThreadsComponent),
          data: { title: 'Threads' },
        },
        {
          path: 'courses',
          loadComponent: () => import('./components/pages/courses/courses.component').then(m => m.CoursesComponent),
          data: { title: 'Courses' },
        },
        {
          path: 'about',
          loadComponent: () => import('./components/pages/about/about.component').then(m => m.AboutComponent),
          data: { title: 'About' },
        },
        {
          path: 'form',
          loadComponent: () => import('./components/pages/test-form/test-form.component').then(m => m.TestFormComponent)
        },

        {
            path: '',
            redirectTo: 'home',
            pathMatch: 'full'
        },

      ]
    },
];
