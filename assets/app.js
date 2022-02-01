(async () => {
  const list = document.createElement('ul');
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

  document.querySelector('#title').innerHTML = data.title;
  document.querySelector('#list').innerHTML = list.innerHTML;

  const terms = document.getElementsByClassName('term-expandable');
  for (let i = 0; i < terms.length; i++) {
    terms[i].addEventListener('click', (e) => {
      e.target.parentElement.querySelector('.children').classList.toggle('active');
      e.target.classList.toggle('term-expanded');
    });
  }
})();
