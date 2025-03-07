Thu Mar  6 06:12:40 PM +0545 2025
---

First day of starting this, lets do a little bit of requirement gathering. Things I want:

- A robust template like feel that can go with anything; documentation, journals, homework ... just the full 9 yards
- Latest Mermaid support
     - The thing I am using now, is an older version of mermaid; which is annoying as there are newer diagrams available.
- Chart.js Support
    - I like the idea of being able to make graphs easily. 
- Vercel's Design Language 
    - Support for icons from [Icons](https://vercel.com/geist/icons)
    - Support for tags / badges [Badges](https://vercel.com/geist/badge)
    - Code Blocks
        - Code Syntax Highlighting.
        - Design: [Code Block Design](https://vercel.com/geist/code-block)
        - Need the Copy Button. 
        - Need the line numbers.
        - Code Snippets [Vercel's Code Snippets](https://vercel.com/geist/snippet)
        - Chroma? <- Package
    - tooltips. [Vercel's ToolTip](https://vercel.com/geist/tooltip)
    - Collapsable things like here [Vercel's Description sectino](https://vercel.com/geist/description)
    - Key Codes like [Vercel's KBD thing](https://vercel.com/geist/keyboard-input)
    - I want a section for the "summary" type shit things we get at the end of like say the math book; that contains the highlights for the entire thing before it.
        - Maybe like this note: [Vercel's Note](https://vercel.com/geist/note)
    - Tables [Vercel's Table](https://vercel.com/geist/table)
    - Tabs [Vercel's Tabs](https://vercel.com/geist/tabs)
    - Typography [Vercel's Typography](https://vercel.com/geist/typography)

- Clickable Check Boxes ( Doesn't Have to preserve state across page refereshes, but i mean just make it checkable atleast )

---

# Understanding Abstract Syntax Trees (AST):
It is a tree representation of structured text. It captures the **hierarchical** structure of the document, making it easier to transform and process. In our case, for markdown, it lets you:

- Recognize blocks like headings,paragraphs,lists, and code blocks.
- Process inline elements like bold,italic and links.

Eg:
```md
~ This:

# Title
**Bold** text
```
```bash
~ Gets converted to:

Root
 ├── Heading(level=1) → "Title"
 ├── Paragraph
     ├── Bold → "Bold"
     ├── Text → " text"
```

# Steps:
1. Tokenize the input (lexing)
    - Split the input by new lines
    - For each line, identify the type of block:
        - Headings, 
        - Lists, 
        - Code blocks (Starts with tripple back-ticks; or indented lines; buut im thinking just make it tripple back-ticks because i want tab to just be like tab yk like just to indent instead of the `>`)

