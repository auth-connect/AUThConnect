
import { CommonModule } from '@angular/common';
import { Component, inject, input, output } from '@angular/core';
import { RouterLink, RouterModule, RouterOutlet } from '@angular/router';
import { NgIconComponent, provideIcons, provideNgIconsConfig } from '@ng-icons/core';
import { featherBook, featherHome, featherMessageCircle,  } from '@ng-icons/feather-icons';
import { bootstrapQuestion } from '@ng-icons/bootstrap-icons'
import { RouteInterface, SidebarService } from '../../../services/sidebar-service/sidebar.service';

@Component({
  selector: 'app-sidebar',
  standalone: true,
  imports: [RouterModule, CommonModule, RouterOutlet, RouterLink, NgIconComponent],
  viewProviders: [
    provideIcons({ featherHome, featherMessageCircle, bootstrapQuestion, featherBook }), 
    provideNgIconsConfig({
      size: '1.7em',
      color: '#d2d3d5',
    }),
  ],
  templateUrl: './sidebar.component.html',
  styleUrl: './sidebar.component.scss',
})
export class SidebarComponent {
  isSidebarCollapsed = input.required<boolean>();
  changeIsSidebarCollapsed = output<boolean>();

  private sidebarService: SidebarService = inject(SidebarService);
  
  public items: RouteInterface[] = this.sidebarService.getRoutes();

  toggleCollapse(): void {
    this.changeIsSidebarCollapsed.emit(!this.isSidebarCollapsed());
  }

  closeSidenav(): void {
    this.changeIsSidebarCollapsed.emit(true);
  }
}
