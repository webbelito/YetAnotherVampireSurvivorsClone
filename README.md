# Game Design Document: Proof of Concept

## 1. Game Overview

### Game Title:  
_Tentative Name_

### Genre:  
Roguelike, Bullet Hell, Action

### Target Platform:  
PC / Web

### Game Summary:  
A survival action game where the player fights off endless waves of enemies, auto-attacking with weapons while collecting power-ups and leveling up to survive as long as possible.

### Unique Selling Points (USP):  
- Procedurally generated enemies.
- Auto-attack mechanics without direct player input.
- Increasing difficulty and scaling enemies.

---

## 2. Game Mechanics

### Core Mechanics:
- **Player Movement**: The player moves using WASD or arrow keys.
- **Auto-Attack**: The player’s weapon attacks enemies automatically at regular intervals.
- **Enemy Behavior**: Enemies constantly spawn and move towards the player, attacking when in range.
- **Health and Death**: The player and enemies have health points. When the player’s health reaches 0, the game is over.

### Weapons and Abilities:

#### Weapons:
1. **Sword Slash**:
   - **Description**: A short-range melee weapon that attacks in an arc in front of the player.
   - **Attack Pattern**: Sweeps in a 90-degree arc, dealing damage to all enemies in range.
   - **Cooldown**: 0.8 seconds.
   - **Upgrade Path**: Can increase the arc size and attack speed with upgrades.

2. **Magic Orb**:
   - **Description**: A ranged weapon that fires a homing projectile at the nearest enemy.
   - **Attack Pattern**: The orb seeks the nearest enemy within a short radius and explodes on impact.
   - **Cooldown**: 1.5 seconds.
   - **Upgrade Path**: Increases orb speed and adds chain lightning effect when upgraded.

#### Abilities:
1. **Dash**:
   - **Description**: The player can quickly dash in the current movement direction, temporarily becoming invincible to avoid enemies or projectiles.
   - **Cooldown**: 3 seconds.
   - **Upgrade Path**: Reduces cooldown and adds damage to enemies in the dash path.

2. **Time Slow**:
   - **Description**: Slows down time for 5 seconds, reducing enemy speed by 50% while allowing the player to move and attack normally.
   - **Cooldown**: 15 seconds.
   - **Upgrade Path**: Increases the duration of the time slow effect or makes enemies take increased damage during the slow.

3. **Energy Shield**:
   - **Description**: Grants a temporary shield that absorbs damage for 8 seconds.
   - **Cooldown**: 20 seconds.
   - **Upgrade Path**: Increases shield duration or adds a reflective damage effect, causing enemies that hit the shield to take damage.

---

## 3. Game World

### Level Design:  
The game is played in a top-down 2D field with no boundaries. Enemies spawn from the edges of the screen and move towards the player.

### Environment:  
- **Background**: Simple static background, such as a flat terrain or dark void.
- **Obstacles**: Could add trees, rocks, or other simple obstacles that the player and enemies must avoid or navigate around.

### Spawn Points:  
- **Enemy Spawn**: Enemies spawn randomly at the edges of the screen, entering from all directions.

---

## 4. Game Progression

### Difficulty Scaling:  
- **Wave System**: Enemies spawn in waves, and each wave introduces more enemies or tougher ones.
- **Time-Based Scaling**: As the player survives longer, enemy health and numbers increase progressively.
- **Experience and Leveling**: Killing enemies grants experience points. Upon leveling up, the player can choose a stat to upgrade or unlock a new ability.

### Win/Loss Conditions:  
- **Victory**: Not applicable (infinite gameplay).
- **Defeat**: The game ends when the player’s health reaches 0.

---

## 5. User Interface (UI)

### HUD (Heads-Up Display):
- **Health Bar**: Shows the player’s current health.
- **Experience Bar**: Shows progress toward the next level.
- **Ability Cooldowns**: Displays the cooldown timers for equipped abilities.
- **Score**: Keeps track of how many enemies have been killed or time survived.

### Menus:
- **Main Menu**: Start, Options, Exit.
- **Pause Menu**: Resume, Options, Exit.

---

## 6. Technical Design

### Graphics:
For a proof of concept, you can stick with basic 2D graphics using simple shapes or sprites:
- **Player**: A colored rectangle or a simple sprite.
- **Enemies**: Different colored shapes representing different enemy types.

### Controls:
- **Movement**: WASD or arrow keys to move the player character.
- **Abilities**: Mapped to specific keys like `Space` for Dash, `Q` for Time Slow, and `E` for Energy Shield.
- **Pause**: `P` key pauses the game.

### Sound:
- **Music**: Optional background music to add atmosphere.
- **Sound Effects**: Simple sound effects for player attacks, enemy hits, and power-up pickups.

### Tools & Libraries:
- **Engine**: Raylib with Go for development.
- **Version Control**: Git.
- **Graphics**: Placeholder graphics made using simple shapes or tools like Aseprite (if sprites are used).

---

## 7. Timeline and Milestones

### Milestones:
- **Week 1**: 
  - Setup game loop and basic player movement.
  - Implement basic enemy spawning and movement.
- **Week 2**: 
  - Add auto-attack system and basic power-ups.
  - Implement basic UI elements (health bar, experience, ability cooldowns).
- **Week 3**: 
  - Add enemy variations and simple difficulty scaling.
  - Test and polish gameplay mechanics.

---

## 8. Conclusion

### Proof of Concept Goal:  
The goal is to develop a working prototype of a survival-based game with basic movement, enemy spawning, auto-attacks, and leveling mechanics. The focus is on core gameplay, with placeholder graphics and minimal polish, to validate the fun and replayability of the game idea.