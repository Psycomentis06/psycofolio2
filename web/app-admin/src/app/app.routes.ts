import { Routes } from '@angular/router';

export const routes: Routes = [
  {
    path: '',
    loadComponent: () => import('./pages/dashboard/dashboard.component').then(c => c.DashboardComponent)
  },
  {
    path: 'performance',
    loadComponent: () => import('./pages/performance/performance.component').then(c => c.PerformanceComponent)
  }
];
