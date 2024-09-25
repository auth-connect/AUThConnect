import { Routes } from '@angular/router';

export const routes: Routes = [
  {
    path: '',
    loadComponent: () => import('./components/shared/layout/layout.component').then(m => m.LayoutComponent),
    children: [
        {
          path: 'home',
          loadComponent: () => import('./components/pages/home/home.component').then(m => m.HomeComponent)
        },
        {
          path: 'threads',
          loadComponent: () => import('./components/pages/threads/threads.component').then(m => m.ThreadsComponent)
        },
        {
          path: 'courses',
          loadComponent: () => import('./components/pages/courses/courses.component').then(m => m.CoursesComponent)
        },
        {
          path: 'about',
          loadComponent: () => import('./components/pages/about/about.component').then(m => m.AboutComponent)
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
    {
      path: 'register',
      loadComponent: () => import('./components/auth/register-page/register-page.component').then(m => m.RegisterPageComponent)
    },
    {
      path: 'login',
      loadComponent: () => import('./components/auth/login-page/login-page.component').then(m => m.LoginPageComponent)
    },
];
