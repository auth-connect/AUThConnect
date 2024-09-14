
import { CommonModule } from '@angular/common';
import { Component, input, output } from '@angular/core';
import { RouterLink, RouterModule, RouterOutlet } from '@angular/router';
import { NgIconComponent, provideIcons, provideNgIconsConfig } from '@ng-icons/core';
import { featherHome, featherMessageCircle } from '@ng-icons/feather-icons';
import { bootstrapQuestion } from '@ng-icons/bootstrap-icons'
@Component({
  selector: 'app-sidebar',
  standalone: true,
  imports: [RouterModule, CommonModule, RouterOutlet, RouterLink, NgIconComponent],
  viewProviders: [
    provideIcons({ featherHome, featherMessageCircle, bootstrapQuestion }), 
    provideNgIconsConfig({
      size: '2em',
      color: '#d2d3d5',
      
    }),
  ],
  templateUrl: './sidebar.component.html',
  styleUrl: './sidebar.component.scss',
})
export class SidebarComponent {
  isSidebarCollapsed = input.required<boolean>();
  changeIsSidebarCollapsed = output<boolean>();
  
  items = [
    {
      routeLink: 'home',
      icon: 'featherHome',
      label: 'Home',
    },
    {
      routeLink: 'threads',
      icon: 'featherMessageCircle',
      label: 'My Threads',
    },
    {
      routeLink: 'about',
      icon: 'bootstrapQuestion',
      label: 'About',
    },
  ];

  toggleCollapse(): void {
    this.changeIsSidebarCollapsed.emit(!this.isSidebarCollapsed());
  }

  closeSidenav(): void {
    this.changeIsSidebarCollapsed.emit(true);
  }
}
