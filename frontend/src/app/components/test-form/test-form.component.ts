import { CommonModule } from '@angular/common';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Component, inject, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';

export interface Payload {
  username: string;
  full_name: string;
  password: string;
  role: string;
  email: string;
}

@Component({
  selector: 'app-test-form',
  standalone: true,
  imports: [
    ReactiveFormsModule,
    CommonModule,
  ],
  templateUrl: './test-form.component.html',
  styleUrls: ['./test-form.component.scss'] // Corrected 'styleUrl' to 'styleUrls'
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
      username: ['', Validators.required],
      full_name: ['', Validators.required],
      email: ['', [Validators.required, Validators.email]], // Added email validator
      password: ['', [Validators.required, Validators.minLength(6)]], // Added minLength validator
      role: ['', Validators.required]
    });
  }

  public onSubmit(form: FormGroup): void {
    if (form.invalid) {
      console.error('Form is invalid');
      return;
    }

    const payload: Payload = {
      username: form.get('username')?.value,
      full_name: form.get('full_name')?.value,
      password: form.get('password')?.value,
      email: form.get('email')?.value,
      role: form.get('role')?.value
    };

    const headers = new HttpHeaders({
      'Content-Type': 'application/json',
      'Accept': 'application/json'
      // Add other custom headers here if necessary
    });

    this.http.post<any>('http://192.168.4.13:8000/users', payload, { headers }).subscribe(
      response => {
        console.log('User Created', response);
        // Optionally, reset the form or provide user feedback here
      },
      error => {
        console.error('Error creating user', error);
        // Handle different error statuses or messages here
      }
    );
  }
}
