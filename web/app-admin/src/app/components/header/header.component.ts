import { Component } from '@angular/core';
import { HeaderGlobalSearchComponent } from '../header-global-search/header-global-search.component';

@Component({
  selector: 'app-header',
  standalone: true,
  imports: [HeaderGlobalSearchComponent],
  templateUrl: './header.component.html',
  styleUrl: './header.component.scss',
})
export class HeaderComponent {}
