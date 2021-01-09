import { Component, OnInit } from '@angular/core';

// import { Ingredient } from '../shared/ingredient.model';

@Component({
  selector: 'app-fileManager',
  templateUrl: './fileManager.component.html',
  styleUrls: ['./fileManager.component.css']
})
export class FileManagerComponent implements OnInit {
  // ingredients: Ingredient[] = [
  //   new Ingredient('Apples', 5),
  //   new Ingredient('Tomatoes', 10),
  // ];

  constructor() { }

  ngOnInit() {
  }

  // onIngredientAdded(ingredient: Ingredient) {
  //   this.ingredients.push(ingredient);
  // }
}
