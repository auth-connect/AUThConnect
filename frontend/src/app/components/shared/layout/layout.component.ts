import { CommonModule } from '@angular/common';
import { Component, computed, HostListener, input, OnInit, signal } from '@angular/core';
import { RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';
import { HeaderComponent } from "../header/header.component";
import { SidebarComponent } from "../sidebar/sidebar.component";

@Component({
  selector: 'app-main-page',
  standalone: true,
  imports: [CommonModule, RouterOutlet, RouterLink, RouterLinkActive, HeaderComponent, SidebarComponent],
  templateUrl: './layout.component.html',
  styleUrl: './layout.component.scss'
})
export class LayoutComponent implements OnInit{

  isSidebarCollapsed = signal<boolean>(false);
  screenWidth = signal<number>(window.innerWidth);

  @HostListener('window:resize')
  onResize() {
    this.screenWidth.set(window.innerWidth);
    if (this.screenWidth() < 768) {
      this.isSidebarCollapsed.set(true);
    }
  }

  ngOnInit(): void {
    this.isSidebarCollapsed.set(this.screenWidth() < 768);
  }

  changeIsSidebarCollapsed(isSidebarCollapsed: boolean): void {
    this.isSidebarCollapsed.set(isSidebarCollapsed);
  }

  
  sizeClass = computed(() => {
    const isSidebarCollapsed = this.isSidebarCollapsed();
    if (isSidebarCollapsed) {
      return '';
    }
    return this.screenWidth() > 768 ? 'w-[calc(100%-16rem)] ml-[16rem]' : 'w-[calc(100%-16rem)] ml-[16rem]';
  });
}
