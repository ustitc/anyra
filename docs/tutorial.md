# Tutorials

- [Set order for a stand-up meeting](#set-order-for-a-stand-up-meeting)
- [Choose a task to be done first](#choose-a-task-to-be-done-first)
- [Cast a fireball](#cast-a-fireball)
- [Generate names for your fantasy race](#generate-names-for-your-fantasy-race)
- [Use Anyra on smartphone](#use-anyra-on-smartphone)

## Set order for a stand-up meeting

You are tired of the same stand-up every day. Spice it up with randomness. Pass your colleague names in `shuffle`:

``` bash
anyra shuffle Oleg Julia Daniele Ahmed Bethel Liu

Ahmed, Julia, Bethel, Liu, Oleg, Daniele
```

You can also store names in a file to reuse it later:

``` bash
touch my_team.txt
echo "Oleg\nJulia\nDaniele\nAhmed\nBethel\nLiu" > my_team.txt
anyra shuffle --file my_team.txt
```

## Choose a task to be done first

Suppose you have such backlog in `todo.txt`:

```
Fix bug SBC-4231
Write documentation for /user endpoint
Delete unused code in order service
Review Mikes merge request
```

And you're frustrated not knowing what to do first. Use `pick` to make that choice for you:

``` bash
anyra pick --file todo.txt

Write documentation for /user endpoint
```

## Cast a fireball

You play DnD as a mage and need to cast a fireball.

- Fireball makes 8d6 fire damage
- You cast it with 5th level spell slot adding 2d6
- Bard buffed you adding one d4
- You have that awesome staff that gives you additional 3 damage 

Build the expression and use `roll`:

``` bash
anyra roll 8d6 + 2d6 + d4 + 3
```

## Generate names for your fantasy race

You have a file `names.txt` with samples of names for your new flying slug race:

```
Korgak
Kendo
Erfas
Losmak
Rezik
Lasdor
```

You ran out of ideas and need to name three new characters. Use `markov` for this:

``` bash
anyra markov --file names.txt --count 3

Lorgak, Los, Korfak
```

## Use Anyra on smartphone

You are away from computer and need one of your scenarios. Anyra is shipped with http server so you can access it from other devices. 


1. Setup server. It can be any VPS/cloud provider or your own computer with [tailscale](https://tailscale.com)

1. Run anyra as a `server`:

``` bash
anyra server

⇨ http server started on [::]:8080
```

2. Open browser on your smartphone and type desired scenario:

```
http://<your_address>:8080/pick?values=Do&values=Delegate&values=Skip&format=plain
```

3. Add address to bookmarks

Api documentation for each scenario can be found [here](./api.md).
