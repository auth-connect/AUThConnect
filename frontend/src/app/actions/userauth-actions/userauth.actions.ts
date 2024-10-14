import { createAction, props } from '@ngrx/store';
import { IUser } from '../../interfaces/user.interface';

export const registerUser = createAction(
  '[Auth] Register User',
  props<{ userData: IUser }>()
);

export const registerSuccess = createAction(
  '[Auth] Register User Success',
  props<{ token: string }>()
);

export const registerFailure = createAction(
  '[Auth] Register User Failure',
  props<{ error: any }>()
);
