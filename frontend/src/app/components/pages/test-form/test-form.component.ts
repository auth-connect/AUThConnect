import { CommonModule } from '@angular/common';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Component, inject, OnInit } from '@angular/core';
import {
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { environment } from '../../../../environment/environment';

export interface Payload {
  name: string;
  password: string;
  email: string;
}

@Component({
  selector: 'app-test-form',
  standalone: true,
  imports: [ReactiveFormsModule, CommonModule],
  templateUrl: './test-form.component.html',
  styleUrls: ['./test-form.component.scss'], // Corrected 'styleUrl' to 'styleUrls'
})
export class TestFormComponent implements OnInit {
  testForm!: FormGroup;
  private http = inject(HttpClient);
  private fb = inject(FormBuilder);

  constructor() {}

  ngOnInit(): void {
    this.initializeForm();
  }

  private initializeForm(): void {
    this.testForm = this.fb.group({
      name: ['', Validators.required],
      email: ['', [Validators.required, Validators.email]], // Added email validator
      password: ['', [Validators.required, Validators.minLength(6)]], // Added minLength validator
    });
  }

  public onSubmit(form: FormGroup): void {
    if (form.invalid) {
      console.error('Form is invalid');
      return;
    }

    const payload: Payload = {
      name: form.get('name')?.value,
      password: form.get('password')?.value,
      email: form.get('email')?.value,
    };

    const headers = new HttpHeaders({
      'Content-Type': 'application/json',
      Accept: 'application/json',
      // Add other custom headers here if necessary
    });

    const apiUrl = `${environment.apiUrl}/v1/users/register`;

    this.http.post<any>(apiUrl, payload, { headers }).subscribe(
      (response) => {
        console.log('User Created', response);
        // Optionally, reset the form or provide user feedback here
      },
      (error) => {
        console.error('Error creating user', error);
        // Handle different error statuses or messages here
      }
    );
  }
}
