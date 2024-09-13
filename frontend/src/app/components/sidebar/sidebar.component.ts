
import { CommonModule } from '@angular/common';
import { Component, input, output } from '@angular/core';
import { RouterLink, RouterModule, RouterOutlet } from '@angular/router';

@Component({
  selector: 'app-sidebar',
  standalone: true,
  imports: [RouterModule, CommonModule, RouterOutlet, RouterLink],
  templateUrl: './sidebar.component.html',
  styleUrl: './sidebar.component.scss',
})
export class SidebarComponent {
  isSidebarCollapsed = input.required<boolean>();
  changeIsSidebarCollapsed = output<boolean>();
  
  items = [
    {
      routeLink: 'dashboard',
      icon: '',
      label: 'Dashboard',
    },
    {
      routeLink: 'products',
      icon: '',
      label: 'Products',
    },
    {
      routeLink: 'pages',
      icon: '',
      label: 'Pages',
    },
    {
      routeLink: 'settings',
      icon: '',
      label: 'Settings',
    },
  ];

  toggleCollapse(): void {
    this.changeIsSidebarCollapsed.emit(!this.isSidebarCollapsed());
  }

  closeSidenav(): void {
    this.changeIsSidebarCollapsed.emit(true);
  }
}
