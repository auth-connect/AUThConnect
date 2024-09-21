import { Component, inject } from '@angular/core';
import { RegisterFormComponent } from "./register-form/register-form.component";
import { Router } from '@angular/router';

@Component({
  selector: 'app-register-page',
  standalone: true,
  imports: [RegisterFormComponent],
  templateUrl: './register-page.component.html',
  styleUrl: './register-page.component.scss'
})
export class RegisterPageComponent {

  router = inject(Router);

  onLoginClick(): void {
    this.router.navigate(['/login']); // Navigate to the login page
  }



}
