import { Component, inject } from '@angular/core';
import { LoginFormComponent } from "./login-form/login-form.component";
import { Router } from '@angular/router';

@Component({
  selector: 'app-login-page',
  standalone: true,
  imports: [LoginFormComponent],
  templateUrl: './login-page.component.html',
  styleUrl: './login-page.component.scss'
})
export class LoginPageComponent {

  router = inject(Router);

  onRegisterClick(): void {
    this.router.navigate(['/register']); // Navigate to the login page
  }

}
