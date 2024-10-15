import { HttpClient, HttpHeaders } from '@angular/common/http';
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
import { environment } from '../../../../../environment/environment';


@Component({
  selector: 'app-register-form',
  standalone: true,
  imports: [NgIconComponent, ReactiveFormsModule],
  viewProviders: [
    provideIcons({ lucideLoaderCircle, lucideGithub }), 
    provideNgIconsConfig({
      size: '1.7em',
      color: 'white',
    }),
  ],
  templateUrl: './register-form.component.html',
  styleUrl: './register-form.component.scss'
})
export class RegisterFormComponent implements OnInit {
  fb = inject(FormBuilder);

  authService = inject(AuthService);
  userAuthFacade = inject(UserAuthFacade);


  isLoading = signal(false);
  isLoadingGithub = signal(false);

  form!: FormGroup;


  constructor(private http: HttpClient) {}
  
  ngOnInit(): void {
    this.initializeForm();
  }

  private initializeForm(){
    this.form = this.fb.nonNullable.group({
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(6)]],
      username: ['', [Validators.required, Validators.minLength(3)]],
    });
  }



  public onSubmit(): void {
    if (this.form.invalid)
      return;
    this.isLoading.set(true);
		setTimeout(() => this.isLoading.set(false), 3000);

    const payload = {
      name: this.form.get('username')?.value,
      password: this.form.get('password')?.value,
      email: this.form.get('email')?.value,
    };

    this.userAuthFacade.registerUser(payload);
  }


}
