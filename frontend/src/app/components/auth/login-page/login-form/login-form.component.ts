import { HttpClient } from '@angular/common/http';
import { Component, inject, OnInit, signal } from '@angular/core';
import {
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { NgIconComponent, provideIcons, provideNgIconsConfig } from '@ng-icons/core';
import { lucideLoaderCircle, lucideGithub } from '@ng-icons/lucide';
import { AuthService } from '../../../../services/auth-service/auth.service';
import { IUser } from '../../../../interfaces/user.interface';
import { UserAuthFacade } from '../../../../facades/userauth-facade/userauth-facade.service';

@Component({
  selector: 'app-login-form',
  standalone: true,
  imports: [NgIconComponent, ReactiveFormsModule],
  viewProviders: [
    provideIcons({ lucideLoaderCircle, lucideGithub }), 
    provideNgIconsConfig({
      size: '1.7em',
      color: 'white',
    }),
  ],
  templateUrl: './login-form.component.html',
  styleUrl: './login-form.component.scss'
})
export class LoginFormComponent implements OnInit {
  fb = inject(FormBuilder);

  authService = inject(AuthService);
  userAuthFacade = inject(UserAuthFacade);


  isLoading = signal(false);
  isLoadingGithub = signal(false);

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

    this.userAuthFacade.loginUser(this.form.getRawValue());
  }

}
