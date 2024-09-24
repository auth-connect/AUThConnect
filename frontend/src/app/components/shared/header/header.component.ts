import { Component, inject, OnInit } from '@angular/core';
import { ActivatedRoute, NavigationEnd, Router } from '@angular/router';
import { filter, map, mergeMap } from 'rxjs';

@Component({
  selector: 'app-header',
  standalone: true,
  imports: [],
  templateUrl: './header.component.html',
  styleUrl: './header.component.scss'
})

export class HeaderComponent implements OnInit {
  pageTitle: string = '';

  private router =  inject(Router);
  private activatedRoute = inject(ActivatedRoute)

  constructor() {}


  ngOnInit(): void {
    // Set the title immediately upon initialization
    this.updateTitle();

    // Update the title on navigation events
    this.router.events
      .pipe(filter(event => event instanceof NavigationEnd))
      .subscribe(() => {
        this.updateTitle();
      });
  }

  private updateTitle() {
    let route = this.activatedRoute;

    while (route.firstChild) {
      route = route.firstChild;
    }

    if (route.snapshot.data && route.snapshot.data['title']) {
      this.pageTitle = route.snapshot.data['title'];
    } else {
      this.pageTitle = '';
    }
  }
}
