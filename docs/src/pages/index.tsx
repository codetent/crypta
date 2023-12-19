import clsx from 'clsx';
import Link from '@docusaurus/Link';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import Layout from '@theme/Layout';
import {FeatureItem, FeatureCard} from '@site/src/components/FeatureCard';
import Heading from '@theme/Heading';

import styles from './index.module.css';

const FeatureList: FeatureItem[] = [
  {
    title: 'Simple',
    Svg: require('@site/static/img/check.svg').default,
    description: (
      <>
        Crypta provides your developer credentials to your build environment as simple as it should be.
      </>
    ),
  },
  {
    title: 'Local',
    Svg: require('@site/static/img/dolly.svg').default,
    description: (
      <>
        Crypta runs fully local without requiring any additional requirements or services.
      </>
    ),
  },
  {
    title: 'Compatible',
    Svg: require('@site/static/img/plug.svg').default,
    description: (
      <>
        Crypta is compatible with CI/CD workflows as well as other secret providers.
      </>
    ),
  },
];

function HomepageHeader() {
  const {siteConfig} = useDocusaurusContext();
  return (
    <header className={clsx('hero hero--primary', styles.heroBanner)}>
      <div className="container">
        <Heading as="h1" className="hero__title">
          {siteConfig.title}
        </Heading>
        <p className="hero__subtitle">{siteConfig.tagline}</p>
        <div className={styles.buttons}>
          <Link
            className="button button--secondary button--lg"
            to="/pages/category/getting-started">
            Getting Started
          </Link>
        </div>
      </div>
    </header>
  );
}

export default function Home(): JSX.Element {
  const {siteConfig} = useDocusaurusContext();
  return (
    <Layout
      title={`Hello from ${siteConfig.title}`}
      description="Description will go into a meta tag in <head />">
      <HomepageHeader />
      <main>
        <section className={styles.features}>
          <div className="container">
            <div className="row">
              {FeatureList.map((props, idx) => (
                <FeatureCard key={idx} {...props} />
              ))}
            </div>
          </div>
        </section>
      </main>
    </Layout>
  );
}
