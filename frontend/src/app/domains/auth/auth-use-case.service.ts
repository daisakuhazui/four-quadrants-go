import { Injectable } from '@angular/core';
import Auth0Client from '@auth0/auth0-spa-js/dist/typings/Auth0Client';
import { from, of, Observable, BehaviorSubject, combineLatest, throwError } from 'rxjs';
import { tap, catchError, concatMap, shareReplay } from 'rxjs/operators';
import createAuth0Client from '@auth0/auth0-spa-js';
import { environment } from '../../../environments/environment';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class AuthUseCaseService {

  auth0Client$ = (from(
    createAuth0Client({
      domain: environment.auth0.domain,
      client_id: environment.auth0.clientId,
      redirect_uri: `${window.location.origin}/callback`
    })
  ) as Observable<Auth0Client>).pipe(
    shareReplay(1),
    catchError(err => throwError(err))
  );

  isAuthenticated$ = this.auth0Client$.pipe(
    concatMap((client: Auth0Client) => from(client.isAuthenticated())),
    tap(res => this.loggedIn = res)
  );

  handleRedirectCallback$ = this.auth0Client$.pipe(
    concatMap((client: Auth0Client) => from(client.handleRedirectCallback()))
  );

  private userProfileSubject$ = new BehaviorSubject<any>(null);

  userProfile$ = this.userProfileSubject$.asObservable();

  loggedIn: boolean = null;

  constructor(private router: Router) { }

  getUser$(options?): Observable<any> {
    return this.auth0Client$.pipe(
      concatMap((client: Auth0Client) => from(client.getUser(options))),
      tap(user => this.userProfileSubject$.next(user)),
    );
  }

  localAuthSetup() {
    const checkAuth$ = this.isAuthenticated$.pipe(
      concatMap((loggedIn: boolean) => {
        if (loggedIn) {
          return this.getUser$();
        }
        return of (this.loggedIn);
      })
    );
    checkAuth$.subscribe((response: { [key: string]: any } | boolean) => {
      this.loggedIn = !!response;
    });
  }

  login(redirectPath: string = '/index') {
    this.auth0Client$.subscribe((client: Auth0Client) => {
      console.log(client);
      client.loginWithRedirect({
        redirect_uri: `${window.location.origin}/callback`,
        appState: {target: redirectPath}
      });
    });
  }

  handleAuthCallback() {
    let targetRoute: string;
    const authComplete$ = this.handleRedirectCallback$.pipe(
      tap(cbRes => {
        targetRoute = cbRes.appState && cbRes.appState.target ? cbRes.appState.target : '/';
      }),
      concatMap(() => {
        return combineLatest(
          this.getUser$(),
          this.isAuthenticated$
        );
      })
    );
    authComplete$.subscribe(([user, loggedIn]) => {
      this.router.navigate([targetRoute]);
    });
  }

  logout() {
    this.auth0Client$.subscribe((client: Auth0Client) => {
      client.logout({
        client_id: environment.auth0.clientId,
        returnTo: `${window.location.origin}`
      });
    });
  }
}
