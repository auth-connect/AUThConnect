import { HttpClient } from '@angular/common/http';
import { Component, inject, OnInit, signal } from '@angular/core';
import {
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { NgIconComponent, provideIcons, provideNgIconsConfig } from '@ng-icons/core';
import { lucideLoader } from '@ng-icons/lucide';
import { AuthService } from '../../../../services/auth-service/auth.service';
import { UserInterface } from '../../../../interfaces/user.interface';
import { AuthFacade } from '../../../../facades/auth-facade/auth-facade.service';

@Component({
  selector: 'app-login-form',
  standalone: true,
  imports: [NgIconComponent, ReactiveFormsModule],
  viewProviders: [
    provideIcons({ lucideLoader }), 
    provideNgIconsConfig({
      size: '1.7em',
      color: '#d2d3d5',
    }),
  ],
  templateUrl: './login-form.component.html',
  styleUrl: './login-form.component.scss'
})
export class LoginFormComponent implements OnInit {
  fb = inject(FormBuilder);

  authService = inject(AuthService);
  authFacade = inject(AuthFacade);


  isLoading = signal(false);

  form!: FormGroup;


  constructor() {}
  
  ngOnInit(): void {
    this.initializeForm();
  }

  initializeForm(){
    this.form = this.fb.nonNullable.group({
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(6)]],
    });
  }



  onSubmit(): void {
    if (this.form.invalid)
      return;
    this.isLoading.set(true);
		setTimeout(() => this.isLoading.set(false), 3000);
    const user: UserInterface = this.form.getRawValue()

    console.log(user);
    this.authFacade.loginUser(user);
  }

}
