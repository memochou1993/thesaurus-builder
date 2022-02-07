Thesaurus Builder
===

Thesaurus Builder is a CLI tool to build a static site of thesaurus.

# DEMO

- [Art & Architecture Thesaurus](https://memochou1993.github.io/aat-simplified/)

# Quick Start

Create a `thesaurus.yaml` file.

```YAML
---
title: Thesaurus
subjects:
  # root is required
  - terms:
      - text: emotion
        preferred: true
      - text: emotions
      - text: 情緒
    notes:
      - text: Refers to a complex phenomena and quality of consciousness, featuring the synthesis or combination of subjective experiences and perceptions, expressive physiological and psychological behaviors, and the excitation or stimulation of the nervous system. Among psychological studies, the concept is associated with ideas on personality formation, rational and irrational thinking, and cognitive motivation.
      - text: 意指意識所呈現的複雜現象及特性。情感涉及了合成或組合主觀經驗與感知、具表達意義的生理與心理行為、以及神經系統的興奮或刺激。在心理學研究中，情感的概念與個性形成、理性與非理性思維、以及認知動機有關。
  # children of root
  - terms:
      - text: love
        preferred: true
      - text: 愛
    parentRelationships:
      - text: emotion
        preferred: true
    notes:
      - text: Emotional and psychological state based on strong affection, loyalty, and benevolence for another arising out of kinship, as in maternal love; arising out of sexual attraction and emotional affinity, as in affection and tenderness felt between lovers; and arising out of respect and admiration, as in the valuation and appreciation among friends.
      - text: 一種衍生自親人之間的強烈關愛、忠誠及善意的情感與心理狀態，如母愛。亦可為衍生自兩性之間在性欲與情感上的吸引力，例如情人之間的情愛與溫柔。此外，亦可能為衍生自尊敬與欽佩之情，例如朋友之間彼此重視與欣賞。
  # children of root
  - terms:
      - text: anger
        preferred: true
      - text: 憤怒
    parentRelationships:
      - text: emotion
        preferred: true
    notes:
      - text: Strong emotional reaction of displeasure or hostility demonstrated by physical reactions, particular facial grimaces and body positions characteristic of action in the autonomic nervous system.
      - text: 意指不滿或敵意所引起的強烈情緒反應。憤怒時，自律神經系統會產生作用，進而引發生理反應，並且使人表現出特有的面部表情與身體姿勢。
```

Build static site with binary.

```BASH
./bin/tb
```

# Usage

## Fields of thesaurus

### title

Title of thesaurus.

### subjects

Subjects of thesaurus.

### subjects[].terms

Terms of subject. There must be and can only be one preferred term in a subject. Non-preferred terms can be multiple.

### subjects[].parentRelationships

Parents of subject. There must be and can only be one preferred parent in a subject. Multiple parents are not implemented. There is no parent in a root subject.

### subjects[].notes

Notes of subject.

## Arguments of CLI

| Argument | Default Value  | Description            |
| -------- | -------------- | ---------------------- |
| -f       | thesaurus.yaml | thesaurus file         |
| -o       | dist           | output directory       |
| -t       | theme          | default                |
| -td      | N/A            | custom theme directory |

## Available Themes

- default
- expanded

## Custom Theme

Copy default theme, and open in browser.

```BASH
cp -R themes/default my-theme
```

Customize styles.

```BASH
body {
  background-color: #000000;
  filter: invert(100%) hue-rotate(30deg);
}
```

Build static site with binary.

```BASH
./bin/tb -td my-theme
```

# Development

Run `main.go` and specify `example.yaml` as a testing file.

```BASH
go run main.go -f example.yaml
```
