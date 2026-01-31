import { Component, signal } from '@angular/core';

@Component({
	selector: 'app-navigation',
	imports: [],
	templateUrl: './navigation.html',
	styleUrl: './navigation.css',
})
export class Navigation {
	isExpanded = signal(false)

	public toggle() {
		this.isExpanded.update(isExpanded => !isExpanded)
	}
}
