import { CommonModule } from '@angular/common';
import { Component, inject, OnInit } from '@angular/core';
import { RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';
import { AuthService } from './services/auth-service/auth.service';
import { AuthFacade } from './facades/auth-facade/auth-facade.service';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, RouterOutlet, RouterLink, RouterLinkActive],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent implements OnInit {

  title = 'frontend';

  authService = inject(AuthService);
  authFacade = inject(AuthFacade);

  ngOnInit(): void {
    this.authFacade.getUser();
  }
}
