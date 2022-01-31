const data = "__DATA__";

const render = (target, item) => {
  const subject = document.createElement('li');
  subject.setAttribute('class','subject');
  target.appendChild(subject);
  item.term.preferredTerms.forEach((v) => {
    const title = document.createElement('div');
    title.innerHTML = v.termText;
    title.setAttribute('class', item.children.length ? 'title title-expandable' : 'title title-expanded');
    subject.appendChild(title);
  });
  item.note.descriptiveNotes.forEach((v) => {
    const note = document.createElement('div');
    note.innerHTML = v.noteText;
    note.setAttribute('class', 'note');
    subject.appendChild(note);
  });
  const children = document.createElement('ul');
  children.setAttribute('class','children');
  subject.appendChild(children);
  item.children.forEach((c) => render(children, c));
};

const list = document.createElement('ul');
render(list, data);
document.querySelector('#list').innerHTML = list.innerHTML;

const titles = document.getElementsByClassName('title-expandable');
for (let i = 0; i < titles.length; i++) {
  titles[i].addEventListener('click', (e) => {
    e.target.parentElement.querySelector('.children').classList.toggle('active');
    e.target.classList.toggle('title-expanded');
  });
}
