import { inject, Injectable } from '@angular/core';
import { LoggerService } from '../logger-service/logger.service';


export interface RouteInterface{
    routeLink: string,
    icon: string,
    label: string,
}

@Injectable({
  providedIn: 'root',
})

export class SidebarService {

  private items: RouteInterface[] = [
    {
      routeLink: '',
      icon: 'featherHome',
      label: 'Home',
    },
    {
      routeLink: 'threads',
      icon: 'featherMessageCircle',
      label: 'My Threads',
    },
    {
      routeLink: 'courses',
      icon: 'featherBook',
      label: 'Courses',
    },
    {
      routeLink: 'about',
      icon: 'bootstrapQuestion',
      label: 'About',
    },
  ];

  constructor() {}

  public getRoutes() {
    return this.items;
  }
  
}