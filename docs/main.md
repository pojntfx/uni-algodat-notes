% Uni Algodat Notes
% Felicitas Pojtinger
% \today
\tableofcontents

# Uni Algodat Notes

## Introduction

### Contributing

These study materials are heavily based on [professor Toenniessen's "Algorithmen und Datenstrukturen" lecture at HdM Stuttgart](https://www.hdm-stuttgart.de/vorlesung_detail?vorlid=5212224).

**Found an error or have a suggestion?** Please open an issue on GitHub ([github.com/pojntfx/uni-algodat-notes](https://github.com/pojntfx/uni-algodat-notes)):

![QR code to source repository](./static/qr.png){ width=150px }

If you like the study materials, a GitHub star is always appreciated :)

### License

![AGPL-3.0 license badge](https://www.gnu.org/graphics/agplv3-155x51.png){ width=128px }

Uni Algodat Notes (c) 2021 Felicitas Pojtinger and contributors

SPDX-License-Identifier: AGPL-3.0
\newpage

## Cheatsheet

I created a hand-written cheatsheet which contains the most important points; you can print it (use a printer with at least 2400 DPI) for use in the exam: [Cheatsheet](./static/cheatsheet.pdf)

## Themen

- Einleitung
  - Zyklus der Informationsverarbeitung
  - Grundbegriffe zu Algodat
  - Eigenschaften von Algorithmen
- Zeitkomplexität
  - Problemgröße des Inputs
  - Schrittzahlfunktionen
  - Testfunktionen
  - Best Case & Worst Case
  - Beispiele
    - Suche im Binärbaum
    - Fibonacci-Sequenz
    - Bubble-Sort
- Applikative Algorithmen
  - Definition Algorithmus
  - Schritte des Auswertung
  - Rekursion
- Sortieralgorithmen
  - Merge-Sort
    - Zeitkomplexität
  - Quicksort
    - Pivot-Element
    - Zeitkomplexität
- Abstrakte Datentypen
  - ADT
  - Stacks
  - Queues
  - Binärbäume
- Binäre Bäume
  - Implementation BinaryTree
  - Suchbäume
    - Definition
    - Zeitkomplexität
- AVL-Bäume
  - Definition
  - Balancing
  - Doppelte Relation
  - Zeitkomplexität

## Paradigm Conversion

I'm too stupid to write proper functional code, so most of the times I "convert" my imperative solutions to functional ones using the following schema I found on the web:

1. Isolate the loop in its own function. Make sure that all captured variables (i.e., variables that are used in the loop, but not declared in the loop) are passed in as parameters.
2. If the loop declares its own counter (e.g., in a for-loop), remove that declaration and make the loop counter another parameter.
3. Replace the loop construct itself:
   1. Replace the loop condition check with an if statement (or if expression, if that’s what your language has).
   2. In the failure branch, return.
   3. Move the rest of the loop code into the success branch.
   4. At the end of the success branch, add a recursive call which passes the modified values of the loop counter and all captured variables; this is equivalent to the jump back to the top of the original loop.
4. At the original site of the loop, put in a call to your new recursive function. This is where you’ll now provide the initial value of the loop counter.
5. Perform whatever other optimizations seem obvious, as the code at this point will be functional, but probably ugly, probably inefficient, and probably unidiomatic.

## Begriffe

- **Algorithmus**: Präzise, endliche Beschreibung eines allgemeinen Verfahrens unter Verwendung ausführbarer, elementarer Schritte
- **Daten**: Von Maschinen verarbeitbare Form von Infos
- **Datenstruktur**: Von Maschinen verarbeitbarer struktureller Rahmen für die Darstellung von Daten
- **Programm**: Formulierung eines oder mehrerer Algorithmen und Datenstrukturen durch Programmiersprache

## Eigenschaften von Algorithmen

- **Terminierend**: Bricht bei jeder (erlaubten) Eingabe nach endlich vielen Schritten ab
- **Determiniertes Ergebnis**: Liefert bei selber (erlaubter) Eingabe immer dasselbe Ergebnis
- **Determinierter Ablauf**: Führt bei selber (erlaubter) Eingabe immer die Schritte in derselben Reihenfolge aus
- **Deterministisch**: Determiniertes Ergebnis & determinierter Ablauf
- **Zeitkomplexität**: Wachstumsverhalten der benötigten Anzahl von Schritten, wenn Problemgröße gegen unendlich strebt

## Zeitkomplexität

1. Ausgangspunkt: Algorithmus $A$
2. Finde ein Maß für die Problemgröße des Inputs
   - **Sortieren einer Liste**: $N = Anzahl\ Elemente\ der\ Liste$
   - **Bäume**: $N = Anzahl\ Knoten$
   - **Textsuche**: $N = Anzahl\ Zeichen\ im\ Text$
3. Finde die Anzahl der signifikanten Schritte von $A$ bei Problemgröße $N$.
4. Bestimme das Wachstumsverhalten der Funktion für $N→∞$
5. Vergleiche Funktion mit $t(N)$ und finde _richtiges_ Wachstumsverhalten.

### Wichtige Wachstumsfunktionen

- **Bäume**: $O(log(n))$
- **Binary Search**: $O(log(n))$
- **Merge Sort**: $O(log(n))$ (Out-of-Place)
- **Quicksort**: $O(log(n))$ (In-Place)
- **Bubble Sort**: $O(n^2)$

### Sortieralgorithmen

- **In-Place**: Es wird (nahezu) kein weiterer Speicher bei der Ausführung gebraucht (i.e. Quicksort)
- **Out-of-Place**: Es wird zusätzlicher Speicher bei der Ausführung gebraucht (i.e. Merge Sort)

## ADTs

### Übersicht

- Abstrakte Datentypen als Vorläufer der Klassen in OOP
- ADT: Mehrfach instanziierbares Software-Modul
- Jede Instanz hat ein Interface
- Syntax für Operatoren: `<op_name>: <arg1> [x arg2...] → <return_type>`

### Stack

- **type** $Stack(T)$
- **sorts**: $T, bool$
- **operations**
  - **new**: $ø → Stack(T)$
  - **push**: $Stack(T) x T → Stack(T)$
  - **pop**: $Stack(T) → Stack(T)$
  - **top**: $Stack(T) → T$
  - **empty**: $Stack(T) → bool$
- **axioms**

  - $pop(push(s, t)) = s$
  - $top(push(s, t)) = t$
  - $empty(new(())) = true$
  - $empty(push(s, t)) = false$

### Queue

- **type** $Queue(T)$
- **sorts** $T, bool$
- **operations**
  - **new**: $ø → Queue(T)$
  - **enter**: $Queue(T) x T → Queue(T)$
  - **leave**: $Queue(T) → Queue(T)$
  - **head**: $Queue(T) → T$
  - **empty**: $Stack(T) → bool$
- **axioms**
  - $empty(new())=true$
  - $head(enter(new(), t))=t$
  - $leave(enter(new(), t))=new()$
  - $leave(new()) = new()$
  - $head(enter(enter(q,s), t)) = head(enter(q,s))$
  - $leave(enter(enter(q,s), t)) = enter(leave(enter(q,s)), t)$

### Binärbaum

- **type** $BinTree(T)$
- **sort** $T, bool$
- **operations**
  - **new**: $ø → BinTree(T)$
  - **bin**: $BinTree(T)xTxBinTree(T) → BinTree(T)$
  - **left**: $BinTree(T) → BinTree(T)$
  - **right**: $BinTree(T) → BinTree(T)$
  - **root**: $BinTree(T) → T$
  - **empty**: $BinTree(T) → bool$
- **axioms**
  - $left(bin(x,t,y))=x$
  - $right(bin(x,t,y))=y$
  - $root(bin(x,t,y))=t$
  - $empty(new())=true$
  - $empty(bin(x,t,y)))=false$
