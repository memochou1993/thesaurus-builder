(async () => {
  const list = document.querySelector('#list');
  const data = await fetch('data.json').then((r) => r.json());
  const render = (target, root) => {
    const subject = document.createElement('li');
    subject.setAttribute('class', 'subject');
    target.appendChild(subject);
    root.subject.term.preferredTerms.forEach((v) => {
      const term = document.createElement('div');
      term.innerHTML = v.termText;
      term.setAttribute('class', root?.children?.length ? 'term term-expandable' : 'term term-expanded');
      subject.appendChild(term);
    });
    root.subject.note?.descriptiveNotes?.forEach((v) => {
      const note = document.createElement('div');
      note.innerHTML = v.noteText;
      note.setAttribute('class', 'note');
      subject.appendChild(note);
    });
    const children = document.createElement('ul');
    children.setAttribute('class', 'children');
    subject.appendChild(children);
    root.children?.forEach((c) => render(children, c));
  };
  render(list, data.root);

  list.addEventListener('click', (e) => {
    if (e.target.classList.contains('term-expandable')) {
      e.target.parentElement.querySelector('.children').classList.toggle('active');
      e.target.classList.toggle('term-expanded');
    }
  });

  const title = document.querySelector('#title');
  title.innerHTML = data.title;
})();
