interface ITree {
  value: number,
  children?: ITree[]
}

const tree: ITree = {
  value: 1,
  children: [
    {
      value: 4,
      children: [
        {
          value: 3
        }
      ]
    },
    {
      value: 2,
      children: [
        {value: 1},
        {value: 2}
      ]
    }
  ]
}

//result = 13

const getTreeValues = (tree: ITree): number[] => {
  const nodes: number[] = []

  nodes.push(tree.value)

  if (tree.children) {
    for (const child of tree.children) {
      nodes.push(...getTreeValues(child))
    }
  }

  return nodes
}


console.log(getTreeValues(tree))




















